import {
  mdiMonitor,
  mdiLock,
  mdiTable,
  mdiListBox, mdiBookEdit,
} from '@mdi/js'

export const superAdminList = [
  {
    to: '/dashboard',
    icon: mdiMonitor,
    label: 'Dashboard'
  },
  {
    to: '/account',
    label: "Account",
    icon: mdiTable
  },
  {
    to:"/showHeritage",
    label: "Heritage",
    icon: mdiListBox,
  },
  {
    to: '/login',
    label: 'Login',
    icon: mdiLock
  },
]

export const adminList = [
  {
    to: '/dashboard',
    icon: mdiMonitor,
    label: 'Dashboard'
  },
  {
    to: '/applyManage',
    label: "ApplyManage",
    icon: mdiTable
  },
  {
    to:"/showHeritage",
    label: "Heritage",
    icon: mdiListBox,
  },
  {
    to: '/login',
    label: 'Login',
    icon: mdiLock
  },
]

export const userList = [
  {
    to: '/dashboard',
    icon: mdiMonitor,
    label: 'Dashboard'
  },
  {
    to:"/showHeritage",
    label: "Heritage",
    icon: mdiListBox,
  },
  {
    to:"/apply",
    label: "Apply",
    icon: mdiBookEdit
  },
  {
    to: '/login',
    label: 'Login',
    icon: mdiLock
  },
]


