<template>
  <div class="form-container">
    <h2>Add Employee</h2>

    <form @submit.prevent="submitForm">
      <div class="form-group">
        <label>Name:</label>
        <input v-model="emp_name" type="text" required />
      </div>

      <div class="form-group">
        <label>Salary:</label>
        <input v-model="salary" type="number" required />
      </div>

      <div class="form-group">
        <label>Department:</label>
        <input v-model="dept_name" type="text" required />
      </div>

      <div class="form-group">
        <label>Programming Language:</label>
        <input v-model="planguage" type="text" required />
      </div>

      <div class="btn-group">
        <button type="submit" class="submit-btn">Submit</button>
        <button type="button" class="clear-btn" @click="clearForm">Clear</button>
      </div>
    </form>

    <button class="back-btn" @click="goBack">⬅ Back</button>
  </div>
</template>

<script setup>
import { ref } from "vue"
import axios from "axios"
import { useRouter } from "vue-router"

const emp_name = ref("")
const salary = ref("")
const dept_name = ref("")
const planguage = ref("")
const router = useRouter()

const submitForm = async () => {
  try {
    await axios.post("http://localhost:8080/employees", {
      emp_name: emp_name.value,
      salary: Number(salary.value),
      dept_name: dept_name.value,
      planguage: planguage.value
    })
    alert("Employee details stored to database ✅")
    clearForm()
  } catch (err) {
    console.error(err)
    alert("Error storing employee ❌")
  }
}

const clearForm = () => {
  emp_name.value = ""
  salary.value = ""
  dept_name.value = ""
  planguage.value = ""
}

const goBack = () => {
  router.push("/")
}
</script>

<style scoped>
.form-container {
  width: 90%;
  max-width: 500px;
  margin: 40px auto;
  background: #fff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 6px;
  font-weight: bold;
}

input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 6px;
}

.btn-group {
  display: flex;
  gap: 10px;
  justify-content: center;
  margin-top: 15px;
}

.submit-btn {
  background: #007bff;
  color: white;
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
}

.clear-btn {
  background: #ffc107;
  color: black;
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
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
