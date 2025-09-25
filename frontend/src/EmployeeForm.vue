<template>
  <div class="form-container">
    <h2>Employee Registration Form</h2>
    <form @submit.prevent="submitForm" class="form-box">
      <div class="form-group">
        <label>Employee Name:</label>
        <input v-model="form.name" type="text" placeholder="Enter Employee name" required />
      </div>

      <div class="form-group">
        <label>Salary (₹):</label>
        <input v-model="form.salary" type="number" placeholder="Enter salary" required />
      </div>

      <div class="form-group">
        <label>Department Name:</label>
        <input v-model="form.dept_name" type="text" placeholder="Enter department" required />
      </div>

      <div class="form-group">
        <label>Developer Language:</label>
        <input v-model="form.dev_lang" type="text" placeholder="Enter programming language" required />
      </div>

      <div class="button-group">
        <button type="submit">Submit</button>
        <button type="button" @click="goToList">Retrieve Employees</button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive } from "vue"
import axios from "axios"
import { useRouter } from "vue-router"

const router = useRouter()

const form = reactive({
  name: "",
  salary: "",
  dept_name: "",
  dev_lang: "",
  test_lang: ""
})

const submitForm = async () => {
  await axios.post("http://localhost:8080/employee", form)
  alert("✅ Employee Added Successfully!")
  Object.keys(form).forEach(key => form[key] = "") // reset form
}

const goToList = () => {
  router.push("/list")
}
</script>

<style scoped>
.form-container {
  width: 400px;
  margin: 50px auto;
  padding: 25px;
  border: 1px solid #ccc;
  border-radius: 10px;
  background: #fafafa;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

h2 {
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

.form-box {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 5px;
  font-weight: bold;
  color: #444;
}

input {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}

button {
  padding: 10px 15px;
  border: none;
  border-radius: 6px;
  background: #007bff;
  color: white;
  cursor: pointer;
  transition: 0.3s;
}

button:hover {
  background: #0056b3;
}

button[type="button"] {
  background: #28a745;
}

button[type="button"]:hover {
  background: #1c7c31;
}
</style>
