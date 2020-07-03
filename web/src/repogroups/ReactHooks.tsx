import { CodeHosts, RepogroupMetadata } from './types'
import { SearchPatternType } from '../../../shared/src/graphql/schema'
import * as React from 'react'
export const reactHooks: RepogroupMetadata = {
    title: 'React Hooks',
    name: 'react-hooks',
    url: '/react-hooks',
    repositories: [
        { name: 'reactjs/rfcs', codehost: CodeHosts.GITHUB },
        { name: 'juliettepretot/hook-into-props', codehost: CodeHosts.GITHUB },
        { name: 'salvoravida/react-universal-hooks', codehost: CodeHosts.GITHUB },
        { name: 'antoinejaussoin/jooks', codehost: CodeHosts.GITHUB },
        { name: 'alibaba/hooks', codehost: CodeHosts.GITHUB },
        { name: 'stevenpersia/captain-hook', codehost: CodeHosts.GITHUB },
        { name: 'chrisjpatty/crooks', codehost: CodeHosts.GITHUB },
        { name: 'latviancoder/hooks-by-example', codehost: CodeHosts.GITHUB },
        { name: 'craig1123/react-recipes', codehost: CodeHosts.GITHUB },
        { name: 'ant-design/sunflower', codehost: CodeHosts.GITHUB },
        { name: '21kb/react-hooks', codehost: CodeHosts.GITHUB },
        { name: 'bdbch/react-devto', codehost: CodeHosts.GITHUB },
        { name: 'bdbch/react-github', codehost: CodeHosts.GITHUB },
        { name: 'bdbch/react-localstorage', codehost: CodeHosts.GITHUB },
        { name: 'elgorditosalsero/react-gtm-hook', codehost: CodeHosts.GITHUB },
        { name: 'avkonst/hookstate', codehost: CodeHosts.GITHUB },
        { name: 'zhixiaoqiang/react-request-hook', codehost: CodeHosts.GITHUB },
        { name: 'kevinwolfdev/formal', codehost: CodeHosts.GITHUB },
        { name: 'alewin/useWorker', codehost: CodeHosts.GITHUB },
        { name: 'MarvelSQ/use-properties-hook', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/component-size', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/document-title', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/document-visibility', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/input-value', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/local-storage', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/network-status', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/online-status', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/window-scroll-position', codehost: CodeHosts.GITHUB },
        { name: 'rehooks/window-size', codehost: CodeHosts.GITHUB },
        { name: 'react-rekindle/use-request', codehost: CodeHosts.GITHUB },
        { name: 'staltz/use-profunctor-state', codehost: CodeHosts.GITHUB },
        { name: 'withvoid/melting-pot', codehost: CodeHosts.GITHUB },
        { name: 'diegohaz/constate', codehost: CodeHosts.GITHUB },
        { name: 'fodau/conuse', codehost: CodeHosts.GITHUB },
        { name: 'ctrlplusb/easy-peasy', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/fetch-suspense', codehost: CodeHosts.GITHUB },
        { name: 'nearform/graphql-hooks', codehost: CodeHosts.GITHUB },
        { name: 'mobxjs/mobx-react-lite', codehost: CodeHosts.GITHUB },
        { name: 'upmostly/modali', codehost: CodeHosts.GITHUB },
        { name: 'momentechnologies/moment-hooks', codehost: CodeHosts.GITHUB },
        { name: 'daniel-dx/nice-hooks', codehost: CodeHosts.GITHUB },
        { name: 'aiven715/promise-hook', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/reactive-react-redux', codehost: CodeHosts.GITHUB },
        { name: 'slorber/react-async-hook', codehost: CodeHosts.GITHUB },
        { name: 'megazazik/react-cached-callback', codehost: CodeHosts.GITHUB },
        { name: 'megazazik/react-context-refs', codehost: CodeHosts.GITHUB },
        { name: 'wellyshen/react-cool-dimensions', codehost: CodeHosts.GITHUB },
        { name: 'wellyshen/react-cool-onclickoutside', codehost: CodeHosts.GITHUB },
        { name: 'wellyshen/react-cool-portal', codehost: CodeHosts.GITHUB },
        { name: 'andy9775/react-declare-form', codehost: CodeHosts.GITHUB },
        { name: 'yeskunall/react-dom-status-hook', codehost: CodeHosts.GITHUB },
        { name: 'shiningjason/react-enhanced-reducer-hook', codehost: CodeHosts.GITHUB },
        { name: 'ilyalesik/react-fetch-hook', codehost: CodeHosts.GITHUB },
        { name: 'csfrequency/react-firebase-hooks', codehost: CodeHosts.GITHUB },
        { name: 'ckedwards/react-form-stateful', codehost: CodeHosts.GITHUB },
        { name: 'kitze/react-hanger', codehost: CodeHosts.GITHUB },
        { name: 'mkosir/react-hook-mighty-mouse', codehost: CodeHosts.GITHUB },
        { name: 'zakariaharti/react-hookedup', codehost: CodeHosts.GITHUB },
        { name: 'ytiurin/react-hook-layout', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/react-hooks-async', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/react-hooks-global-state', codehost: CodeHosts.GITHUB },
        { name: 'use-hooks/react-hooks-image-size', codehost: CodeHosts.GITHUB },
        { name: 'beizhedenglong/react-hooks-lib', codehost: CodeHosts.GITHUB },
        { name: 'kmkzt/react-hooks-svgdrawing', codehost: CodeHosts.GITHUB },
        { name: 'shibe97/react-hooks-use-modal', codehost: CodeHosts.GITHUB },
        { name: 'kmkzt/react-hooks-visible', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/react-hooks-worker', codehost: CodeHosts.GITHUB },
        { name: 'JohannesKlauss/react-hotkeys-hook', codehost: CodeHosts.GITHUB },
        { name: 'sin/react-immer-hooks', codehost: CodeHosts.GITHUB },
        { name: 'marceloadsj/react-indicative-hooks', codehost: CodeHosts.GITHUB },
        { name: 'AvraamMavridis/react-intersection-visible-hook', codehost: CodeHosts.GITHUB },
        { name: 'lessmess-dev/react-media-hook', codehost: CodeHosts.GITHUB },
        { name: 'lordgiotto/react-metatags-hook', codehost: CodeHosts.GITHUB },
        { name: 'mamal72/react-optimistic-ui-hook', codehost: CodeHosts.GITHUB },
        { name: 'RyanFitzgerald/react-page-name', codehost: CodeHosts.GITHUB },
        { name: 'vardius/react-peer-data', codehost: CodeHosts.GITHUB },
        { name: 'dispix/react-pirate', codehost: CodeHosts.GITHUB },
        { name: 'kalcifer/react-powerhooks', codehost: CodeHosts.GITHUB },
        { name: 'moxystudio/react-promiseful', codehost: CodeHosts.GITHUB },
        { name: 'hupe1980/react-recaptcha-hook', codehost: CodeHosts.GITHUB },
        { name: 'schettino/react-request-hook', codehost: CodeHosts.GITHUB },
        { name: 'inmagik/react-rocketjump', codehost: CodeHosts.GITHUB },
        { name: 'hupe1980/react-script-hook', codehost: CodeHosts.GITHUB },
        { name: 'Andarist/react-selector-hooks', codehost: CodeHosts.GITHUB },
        { name: 'MikeyParton/react-speech-kit', codehost: CodeHosts.GITHUB },
        { name: 'mcclayton/react-state-patterns', codehost: CodeHosts.GITHUB },
        { name: 'FormidableLabs/react-swipeable', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/react-tracked', codehost: CodeHosts.GITHUB },
        { name: 'j-a-y-h/react-uniformed', codehost: CodeHosts.GITHUB },
        { name: 'RyanRoll/react-use-api', codehost: CodeHosts.GITHUB },
        { name: 'gregnb/react-use-calendar', codehost: CodeHosts.GITHUB },
        { name: 'danoc/react-use-clipboard', codehost: CodeHosts.GITHUB },
        { name: 'smmoosavi/react-use-data-loader', codehost: CodeHosts.GITHUB },
        { name: 'JohannesKlauss/react-use-fetch-factory', codehost: CodeHosts.GITHUB },
        { name: 'grug/react-use-fetch-with-redux', codehost: CodeHosts.GITHUB },
        { name: 'wsmd/react-use-form-state', codehost: CodeHosts.GITHUB },
        { name: 'Yaska/react-use-id-hook', codehost: CodeHosts.GITHUB },
        { name: 'kigiri/react-use-idb', codehost: CodeHosts.GITHUB },
        { name: 'CurationCorp/react-use-infinite-loader', codehost: CodeHosts.GITHUB },
        { name: 'robcalcroft/react-use-input', codehost: CodeHosts.GITHUB },
        { name: 'robcalcroft/react-use-lazy-load-image', codehost: CodeHosts.GITHUB },
        { name: 'intercaetera/react-use-message-bar', codehost: CodeHosts.GITHUB },
        { name: 'wowlusitong/react-use-modal', codehost: CodeHosts.GITHUB },
        { name: 'zhangkaiyulw/react-use-path', codehost: CodeHosts.GITHUB },
        { name: 'neo/react-use-scroll-position', codehost: CodeHosts.GITHUB },
        { name: 'lessmess-dev/react-use-trigger', codehost: CodeHosts.GITHUB },
        { name: 'perlin-network/react-use-wavelet', codehost: CodeHosts.GITHUB },
        { name: 'streamich/react-use', codehost: CodeHosts.GITHUB },
        { name: 'GeDiez/react-use-formless', codehost: CodeHosts.GITHUB },
        { name: 'venil7/react-usemiddleware', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/react-useportal', codehost: CodeHosts.GITHUB },
        { name: 'vardius/react-user-media', codehost: CodeHosts.GITHUB },
        { name: 'f/react-wait', codehost: CodeHosts.GITHUB },
        { name: 'AvraamMavridis/react-window-communication-hook', codehost: CodeHosts.GITHUB },
        { name: 'yesmeck/react-with-hooks', codehost: CodeHosts.GITHUB },
        { name: 'mfrachet/reaktion', codehost: CodeHosts.GITHUB },
        { name: 'iusehooks/redhooks', codehost: CodeHosts.GITHUB },
        { name: 'ianobermiller/redux-react-hook', codehost: CodeHosts.GITHUB },
        { name: 'regionjs/region-core', codehost: CodeHosts.GITHUB },
        { name: 'imbhargav5/rehooks-visibility-sensor', codehost: CodeHosts.GITHUB },
        { name: 'pedronasser/resynced', codehost: CodeHosts.GITHUB },
        { name: 'brn/rrh', codehost: CodeHosts.GITHUB },
        { name: 'LeetCode-OpenSource/rxjs-hooks', codehost: CodeHosts.GITHUB },
        { name: 'dejorrit/scroll-data-hook', codehost: CodeHosts.GITHUB },
        { name: 'style-hook/style-hook', codehost: CodeHosts.GITHUB },
        { name: 'vercel/swr', codehost: CodeHosts.GITHUB },
        { name: 'jaredpalmer/the-platform', codehost: CodeHosts.GITHUB },
        { name: 'danieldelcore/trousers', codehost: CodeHosts.GITHUB },
        { name: 'mauricedb/use-abortable-fetch', codehost: CodeHosts.GITHUB },
        { name: 'awmleer/use-action', codehost: CodeHosts.GITHUB },
        { name: 'awmleer/use-async-memo', codehost: CodeHosts.GITHUB },
        { name: 'lowewenzel/use-autocomplete', codehost: CodeHosts.GITHUB },
        { name: 'sergey-s/use-axios-react', codehost: CodeHosts.GITHUB },
        { name: 'zcallan/use-browser-history', codehost: CodeHosts.GITHUB },
        { name: 'samjbmason/use-cart', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-clippy', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/use-context-selector', codehost: CodeHosts.GITHUB },
        { name: 'oktaysenkan/use-countries', codehost: CodeHosts.GITHUB },
        { name: 'xnimorz/use-debounce', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-deep-compare', codehost: CodeHosts.GITHUB },
        { name: 'kentcdodds/use-deep-compare-effect', codehost: CodeHosts.GITHUB },
        { name: 'gregnb/use-detect-print', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-dimensions', codehost: CodeHosts.GITHUB },
        { name: 'zattoo/use-double-click', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-events', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-force-update', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-hotkeys', codehost: CodeHosts.GITHUB },
        { name: 'therealparmesh/use-hovering', codehost: CodeHosts.GITHUB },
        { name: 'ava/use-http', codehost: CodeHosts.GITHUB },
        { name: 'immerjs/use-immer', codehost: CodeHosts.GITHUB },
        { name: 'immerjs/immer', codehost: CodeHosts.GITHUB },
        { name: 'neighborhood999/use-input-file', codehost: CodeHosts.GITHUB },
        { name: 'helderburato/use-is-mounted-ref', codehost: CodeHosts.GITHUB },
        { name: 'davidicus/use-lang-direction', codehost: CodeHosts.GITHUB },
        { name: 'streamich/use-media', codehost: CodeHosts.GITHUB },
        { name: 'dimitrinicolas/use-mouse-action', codehost: CodeHosts.GITHUB },
        { name: 'jschloer/use-multiselect', codehost: CodeHosts.GITHUB },
        { name: 'wellyshen/use-places-autocomplete', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-popper', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/use-react-modal', codehost: CodeHosts.GITHUB },
        { name: 'CharlesStover/use-react-router', codehost: CodeHosts.GITHUB },
        { name: 'tedstoychev/use-reactive-state', codehost: CodeHosts.GITHUB },
        { name: 'dai-shi/use-reducer-async', codehost: CodeHosts.GITHUB },
        { name: 'flepretre/use-redux', codehost: CodeHosts.GITHUB },
        { name: 'tudorgergely/use-scroll-to-bottom', codehost: CodeHosts.GITHUB },
        { name: 'sandiiarov/use-simple-undo', codehost: CodeHosts.GITHUB },
        { name: 'mfrachet/use-socketio', codehost: CodeHosts.GITHUB },
        { name: 'iamgyz/use-socket.io-client', codehost: CodeHosts.GITHUB },
        { name: 'kmoskwiak/useSSE', codehost: CodeHosts.GITHUB },
        { name: 'alex-cory/use-ssr', codehost: CodeHosts.GITHUB },
        { name: 'haydn/use-state-snapshots', codehost: CodeHosts.GITHUB },
        { name: 'philipp-spiess/use-substate', codehost: CodeHosts.GITHUB },
        { name: 'octet-stream/use-suspender', codehost: CodeHosts.GITHUB },
        { name: 'streamich/use-t', codehost: CodeHosts.GITHUB },
        { name: 'homerchen19/use-undo', codehost: CodeHosts.GITHUB },
        { name: 'donavon/use-dark-mode', codehost: CodeHosts.GITHUB },
        { name: 'KATT/use-is-typing', codehost: CodeHosts.GITHUB },
        { name: 'pranesh239/use-key-capture', codehost: CodeHosts.GITHUB },
        { name: 'tranbathanhtung/usePosition', codehost: CodeHosts.GITHUB },
        { name: 'Tweries/useReducerWithLocalStorage', codehost: CodeHosts.GITHUB },
        { name: 'pankod/react-hooks-screen-type', codehost: CodeHosts.GITHUB },
        { name: 'Purii/react-use-scrollspy', codehost: CodeHosts.GITHUB },
        { name: 'JCofman/react-hook-use-service-worker', codehost: CodeHosts.GITHUB },
        { name: 'bboydflo/use-value-after', codehost: CodeHosts.GITHUB },
        { name: 'wednesday-solutions/react-screentype-hook', codehost: CodeHosts.GITHUB },
    ],
    description: 'Examples of useState for ReactHooks.',
    examples: [
        {
            title: 'useState imports regex search:',
            exampleQuery: <>import [^;]+useState[^;]+ from 'react'</>,
            rawQuery: "import [^;]+useState[^;]+ from 'react'",
            patternType: SearchPatternType.regexp,
        },
        {
            title: 'useState with objects as input parameters structural search:',
            exampleQuery: (
                <>
                    {'useState({:[string]})'} <span className="repogroup-page__keyword-text">count:</span>1000
                </>
            ),
            rawQuery: 'useState({:[string]}) count:1000',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'useState with arrays as input parameters structural search:',
            exampleQuery: (
                <>
                    useState([:[string]]) <span className="repogroup-page__keyword-text">count:</span>1000
                </>
            ),
            rawQuery: 'useState([:[string]]) count:1000',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'useState with any type of input parameters structural search:',
            exampleQuery: (
                <>
                    useState(:[string]) <span className="repogroup-page__keyword-text">count:</span>1000
                </>
            ),
            rawQuery: 'useState(:[string]) count:1000',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'useState with any type of input parameters structural search for only typescript files:',
            exampleQuery: (
                <>
                    useState(:[string]) <span className="repogroup-page__keyword-text">count:</span>1000{' '}
                    <span className="repogroup-page__keyword-text">lang:</span>typescript
                </>
            ),
            rawQuery: 'useState(:[string]) count:1000 lang:typescript',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'useState with exactly two input params for structural search should return a lot fewer results:',
            exampleQuery: (
                <>
                    useState([:[1.], :[2.]]) <span className="repogroup-page__keyword-text">count:</span>1000
                </>
            ),
            rawQuery: 'useState([:[1.], :[2.]]) count:1000',
            patternType: SearchPatternType.structural,
        },
        {
            title: 'useState with two or more params in a specific file with structural search:',
            exampleQuery: (
                <>
                    useState([:[1], :[2]]) <span className="repogroup-page__keyword-text">count:</span>1000{' '}
                    <span className="repogroup-page__keyword-text">file:</span>
                    docs/src/pages/components/transfer-list/TransferList.js
                </>
            ),
            rawQuery: 'useState([:[1], :[2]]) count:1000 file:docs/src/pages/components/transfer-list/TransferList.js',
            patternType: SearchPatternType.structural,
        },
    ],
    homepageDescription: 'Examples of useState for ReactHooks.',
    homepageIcon: 'https://cdn4.iconfinder.com/data/icons/logos-3/600/React.js_logo-512.png',
}
