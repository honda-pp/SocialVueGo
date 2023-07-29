import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import UserList from '../views/UserList.vue';

const routes = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/user-list',
    component: UserList,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
