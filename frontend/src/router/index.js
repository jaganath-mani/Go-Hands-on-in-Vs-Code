import { createRouter, createWebHistory } from "vue-router"
import HomePage from "../views/HomePage.vue"
import EmployeeForm from "../components/EmployeeForm.vue"
import EmployeeList from "../components/EmployeeList.vue"

const routes = [
  { path: "/", name: "Home", component: HomePage },
  { path: "/add", name: "AddEmployee", component: EmployeeForm },
  { path: "/list", name: "EmployeeList", component: EmployeeList },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
