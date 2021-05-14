import Vue from 'vue'
import VueRouter from 'vue-router'
import Editor from './components/Editor.vue'
import Settings from './components/Settings.vue'

Vue.use(VueRouter)

const routes = [
  { component: Editor, name: 'Editor', path: '/editor' },
  { component: Settings, name: 'Settings', path: '/settings' },
]

const router = new VueRouter({
  mode: 'abstract', // mode must be set to 'abstract'
  routes,
})

export default router
