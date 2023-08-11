import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Login from '../views/Login.vue';
import UserList from '../views/UserList.vue';
import TweetList from '../views/TweetList.vue';
import UserProfile from '../views/UserProfile.vue';

const routes = [
  {
    path: '/',
    component: Home,
    meta: { requiresAuth: true },
  },
  {
    path: '/login',
    component: Login,
    meta: { requiresAuth: false },
  },
  {
    path: '/user-list',
    component: UserList,
    meta: { requiresAuth: true },
  },
  {
    path: '/tweet-list',
    name: 'TweetList',
    component: TweetList,
    meta: { requiresAuth: true },
  },
  {
    path: '/:userID',
    name: 'UserProfile',
    component: UserProfile,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('userID') != null;
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  if (requiresAuth && !isLoggedIn) {
    next('/login');
  } else if (!requiresAuth && isLoggedIn) {
    next('/');
  } else {
    next();
  }
});

export default router;
