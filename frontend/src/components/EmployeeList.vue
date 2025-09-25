<template>
  <div class="list-container">
    <h2>Employee List</h2>

    <div class="card">
      <table>
        <thead>
          <tr>
            <th>Emp ID</th>
            <th>Name</th>
            <th>Department</th>
            <th>Programming Language</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="emp in employees" :key="emp.emp_id">
            <td>{{ emp.emp_id }}</td>
            <td>{{ emp.emp_name }}</td>
            <td>{{ emp.dept_name }}</td>
            <td><span class="dev-lang">{{ emp.planguage }}</span></td>
          </tr>
        </tbody>
      </table>
    </div>

    <button class="back-btn" @click="goBack">⬅ Back</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import axios from "axios"
import { useRouter } from "vue-router"

const employees = ref([])
const router = useRouter()

onMounted(async () => {
  const res = await axios.get("http://localhost:8080/employees")
  employees.value = res.data
})

const goBack = () => {
  router.push("/")
}
</script>

<style scoped>
.list-container {
  width: 90%;
  max-width: 950px;
  margin: 40px auto;
  text-align: center;
}

h2 {
  margin-bottom: 20px;
  color: #333;
}

.card {
  background: #ffffff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

thead {
  background: #007bff;
  color: #fff;
}

th, td {
  padding: 12px 15px;
  border-bottom: 1px solid #ddd;
  text-align: center;
}

tbody tr:hover {
  background-color: #f5f5f5;
}

.dev-lang {
  color: #39ff14;
  font-weight: bold;
}

.back-btn {
  margin-top: 20px;
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
  background: #28a745;
  color: white;
  cursor: pointer;
}
</style>
