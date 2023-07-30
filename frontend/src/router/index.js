import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import UserList from '../views/UserList.vue';
import TweetList from '../views/TweetList.vue';

const routes = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/user-list',
    component: UserList,
  },
  {
    path: '/tweet-list',
    name: 'TweetList',
    component: TweetList,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
