import {createRouter, createWebHistory} from 'vue-router';
import MainLayout from '../layouts/MainLayout.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: MainLayout,
      children: [
        {
          path: '',
          name: 'MainPage',
          component: () => import('../views/MainPage.vue')
        },
        {
          // E-posta linkinden token ile gelindiğinde MainLayout içinde
          // ResetPasswordDialog otomatik olarak açılır (MainLayout'taki route watcher)
          path: 'reset-password',
          name: 'ResetPassword',
          component: () => import('../views/MainPage.vue')
        }
      ]
    }
  ],
});

export default router;
