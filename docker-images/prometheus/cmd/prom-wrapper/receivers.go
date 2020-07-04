package main

import (
	"fmt"
	"net/url"

	amconfig "github.com/prometheus/alertmanager/config"
	commoncfg "github.com/prometheus/common/config"
	"github.com/sourcegraph/sourcegraph/schema"
)

const (
	alertmanagerNoopReceiver     = "src-noop-receiver"
	alertmanagerWarningReceiver  = "src-warning-receiver"
	alertmanagerCriticalReceiver = "src-critical-receiver"
)

var (
	// alertmanager notification template reference: https://prometheus.io/docs/alerting/latest/notifications
	alertSolutionsTemplate    = `https://docs.sourcegraph.com/admin/observability/alert_solutions#{{ .CommonLabels.service_name }}-{{ .CommonLabels.name | reReplaceAll "(_low|_high)$" "" | reReplaceAll "_" "-" }}`
	notificationTitleTemplate = "[{{ .CommonLabels.level | toUpper }}] {{ .CommonLabels.description }}"
	notificationBodyTemplate  = fmt.Sprintf(`{{ .CommonLabels.level | title }} alert '{{ .CommonLabels.name }}' is firing for service '{{ .CommonLabels.service_name }}'{{ range .Alerts }} {{ .Labels.instance }}{{ end }}.

For possible solutions, see %s`, alertSolutionsTemplate)
)

// newReceivers converts the given alerts from Sourcegraph site configuration into Alertmanager receivers.
// Each alert level has a receiver, which has configuration for all channels for that level.
func newReceivers(newAlerts []*schema.ObservabilityAlerts, newProblem func(error)) []*amconfig.Receiver {
	var (
		warningReceiver  = &amconfig.Receiver{Name: alertmanagerWarningReceiver}
		criticalReceiver = &amconfig.Receiver{Name: alertmanagerCriticalReceiver}
	)

	for i, alert := range newAlerts {
		var receiver *amconfig.Receiver
		if alert.Level == "critical" {
			receiver = criticalReceiver
		} else {
			receiver = warningReceiver
		}

		notifierConfig := amconfig.NotifierConfig{
			VSendResolved: alert.SendResolved,
		}

		notifier := alert.Notifier
		switch {
		case notifier.Email != nil:
			receiver.EmailConfigs = append(receiver.EmailConfigs, &amconfig.EmailConfig{
				To: notifier.Email.Address,

				Headers: map[string]string{
					"subject": notificationTitleTemplate,
				},
				HTML: fmt.Sprintf(`<body>%s</body>`, notificationBodyTemplate),
				// SMTP configuration is applied globally by changeSMTP

				NotifierConfig: notifierConfig,
			})
		case notifier.Opsgenie != nil:
			var apiURL *amconfig.URL
			if notifier.Opsgenie.ApiUrl != "" {
				url, err := url.Parse(notifier.Opsgenie.ApiUrl)
				if err != nil {
					newProblem(fmt.Errorf("failed to apply notifier %d: %w", i, err))
					continue
				}
				apiURL = &amconfig.URL{URL: url}
			}
			receiver.OpsGenieConfigs = append(receiver.OpsGenieConfigs, &amconfig.OpsGenieConfig{
				APIKey: amconfig.Secret(notifier.Opsgenie.ApiKey),
				APIURL: apiURL,

				Message:     notificationTitleTemplate,
				Description: notificationBodyTemplate,

				NotifierConfig: notifierConfig,
			})
		case notifier.Pagerduty != nil:
			var apiURL *amconfig.URL
			if notifier.Pagerduty.ApiUrl != "" {
				url, err := url.Parse(notifier.Pagerduty.ApiUrl)
				if err != nil {
					newProblem(fmt.Errorf("failed to apply notifier %d: %w", i, err))
					continue
				}
				apiURL = &amconfig.URL{URL: url}
			}
			receiver.PagerdutyConfigs = append(receiver.PagerdutyConfigs, &amconfig.PagerdutyConfig{
				RoutingKey: amconfig.Secret(notifier.Pagerduty.RoutingKey),
				Severity:   notifier.Pagerduty.Severity,
				URL:        apiURL,

				Description: notificationTitleTemplate,
				Links: []amconfig.PagerdutyLink{{
					Text: "Alert solutions",
					Href: alertSolutionsTemplate,
				}},

				NotifierConfig: notifierConfig,
			})
		case notifier.Slack != nil:
			url, err := url.Parse(notifier.Slack.Url)
			if err != nil {
				newProblem(fmt.Errorf("failed to apply notifier %d: %w", i, err))
				continue
			}
			if notifier.Slack.Username != "" {
				notifier.Slack.Username = "Sourcegraph Alerts"
			}
			receiver.SlackConfigs = append(receiver.SlackConfigs, &amconfig.SlackConfig{
				APIURL:    &amconfig.SecretURL{URL: url},
				Username:  notifier.Slack.Username,
				Channel:   notifier.Slack.Recipient,
				IconEmoji: notifier.Slack.Icon_emoji,
				IconURL:   notifier.Slack.Icon_url,

				Title: notificationTitleTemplate,
				Text:  notificationBodyTemplate,

				NotifierConfig: notifierConfig,
			})
		case notifier.Webhook != nil:
			url, err := url.Parse(notifier.Webhook.Url)
			if err != nil {
				newProblem(fmt.Errorf("failed to apply notifier %d: %w", i, err))
				continue
			}
			receiver.WebhookConfigs = append(receiver.WebhookConfigs, &amconfig.WebhookConfig{
				URL: &amconfig.URL{URL: url},
				HTTPConfig: &commoncfg.HTTPClientConfig{
					BasicAuth: &commoncfg.BasicAuth{
						Username: notifier.Webhook.Username,
						Password: commoncfg.Secret(notifier.Webhook.Password),
					},
					BearerToken: commoncfg.Secret(notifier.Webhook.BearerToken),
				},

				NotifierConfig: notifierConfig,
			})
		default:
			newProblem(fmt.Errorf("failed to apply notifier %d: no configuration found", i))
		}
	}

	return []*amconfig.Receiver{warningReceiver, criticalReceiver}
}