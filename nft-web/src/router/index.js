import { createRouter, createWebHistory } from 'vue-router';
import Markets from '../views/Markets/index.vue';

const routes = [
    {
      path: '/',
      redirect: '/Markets'
    },
    {
      path: '/Markets',
      name: 'Markets',
      component: Markets
    },
    {
      path: '/Collection',
      name: 'Collection',
      component: () => import(/* webpackChunkName: "Setting" */ '@/views/Collection/index.vue')
    },
    {
      path: '/CollectionDetail',
      name: 'CollectionDetail',
      component: () => import(/* webpackChunkName: "Setting" */ '@/views/Collection/detail.vue')
    },
    {
      path: '/ViewMore',
      name: 'ViewMore',
      component: () => import(/* webpackChunkName: "Setting" */ '@/views/ViewMore/index.vue')
    },
    {
      path: '/Setting',
      name: 'Setting',
      component: () => import(/* webpackChunkName: "Setting" */ '@/views/Setting/index.vue')
    }
  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  // }
  ],

  router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
  });

export default router;
