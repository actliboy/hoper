import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";
import "@/plugin/axios";
import {
  Col,
  Row,
  NavBar,
  Tabbar,
  TabbarItem,
  Tab,
  Tabs,
  Icon,
  List,
  Skeleton,
  Cell,
  CellGroup,
  Toast,
  Notify,
  Form,
  Field,
  Button
} from "vant";

createApp(App)
  .use(store)
  .use(router)
  .use(Col)
  .use(Row)
  .use(NavBar)
  .use(Tabbar)
  .use(TabbarItem)
  .use(Tab)
  .use(Tabs)
  .use(Icon)
  .use(List)
  .use(Skeleton)
  .use(Cell)
  .use(CellGroup)
  .use(Toast)
  .use(Notify)
  .use(Form)
  .use(Field)
  .use(Button)
  .mount("#app");
