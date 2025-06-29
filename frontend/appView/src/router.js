import { createRouter, createWebHistory } from 'vue-router'

import ProfilesPage from './pages/ProfilesPage.vue'
import StatsPage from './pages/StatsPage.vue'
import SettingsPage from './pages/SettingsPage.vue'
import AccountPage from './pages/AccountPage.vue'

const routes = [
  { path: '/', redirect: '/profiles' },
  { path: '/profiles', component: ProfilesPage },
  { path: '/stats', component: StatsPage },
  { path: '/settings', component: SettingsPage },
  { path: '/account', component: AccountPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router 