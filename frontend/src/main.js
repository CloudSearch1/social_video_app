import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/main.css'

// 导入Vant组件
import { 
  Button, 
  Tabbar, 
  TabbarItem, 
  NavBar, 
  Icon, 
  Swipe, 
  SwipeItem,
  Cell,
  CellGroup,
  Image as VanImage,
  Lazyload,
  PullRefresh,
  List,
  Loading,
  Toast,
  Dialog
} from 'vant'
import 'vant/lib/index.css'

const app = createApp(App)

// 注册Vant组件
app.use(Button)
app.use(Tabbar)
app.use(TabbarItem)
app.use(NavBar)
app.use(Icon)
app.use(Swipe)
app.use(SwipeItem)
app.use(Cell)
app.use(CellGroup)
app.use(VanImage)
app.use(Lazyload)
app.use(PullRefresh)
app.use(List)
app.use(Loading)
app.use(Toast)
app.use(Dialog)

app.use(createPinia())
app.use(router)

app.mount('#app')