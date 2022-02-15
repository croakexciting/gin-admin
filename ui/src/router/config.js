import TabsView from '@/layouts/tabs/TabsView'
import BlankView from '@/layouts/BlankView'
import PageView from '@/layouts/PageView'

// 路由配置
const options = {
  routes: [
    {
      path: '/login',
      name: '登录页',
      component: () => import('@/pages/login')
    },
    {
      path: '*',
      name: '404',
      component: () => import('@/pages/exception/404'),
    },
    {
      path: '/403',
      name: '403',
      component: () => import('@/pages/exception/403'),
    },
    {
      path: '/',
      name: '首页',
      component: TabsView,
      redirect: '/login',
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          meta: {
            icon: 'dashboard'
          },
          component: BlankView,
          children: [
            {
              path: 'workplace',
              name: '工作台',
              meta: {
                page: {
                  closable: false
                }
              },
              component: () => import('@/pages/dashboard/workplace'),
            },
          ]
        },
        {
          path: 'demo',
          name: '示例',
          meta: {
            icon: 'table',
          },
          component: PageView,
          children: [
            {
              path: 'query',
              name: '示例列表',
              component: () => import('@/pages/mydemo'),
            },
          ],
        },
        {
          path: 'user',
          name: '用户管理',
          meta: {
            icon: 'table',
          },
          component: PageView,
          children: [
            {
              path: 'query',
              name: '用户列表',
              component: () => import('@/pages/user'),
            },
          ],
        },
      ]
    },
  ]
}

export default options
