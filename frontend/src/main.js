import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faTerminal, faCog } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import * as Wails from '@wailsapp/runtime';
import router from './router'

Vue.config.productionTip = false;
Vue.config.devtools = true;

// Font awesome.
library.add(faTerminal, faCog)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Wails.Init(() => {
	new Vue({
		render: h => h(App),
		router,
    mounted() {
      this.$router.replace('/')
			let root = this.myprop || '/editor'
			this.$router.push({ path: root, query: { name: 'Editor' } });
    },
	}).$mount('#app');
});
