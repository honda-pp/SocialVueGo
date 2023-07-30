import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

const app = createApp(App);

app.use(router);
app.use(store);

document.title = 'SocialVueGo';

store.dispatch('checkLoginStatus').then(() => {
  app.mount('#app');
});
