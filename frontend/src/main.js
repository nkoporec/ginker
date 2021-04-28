import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faTerminal, faCog } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

// Font awesome.
library.add(faTerminal, faCog)

Vue.component('font-awesome-icon', FontAwesomeIcon)


Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
