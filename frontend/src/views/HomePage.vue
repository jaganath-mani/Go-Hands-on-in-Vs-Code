<template>
  <div class="home-container">
    <div class="overlay">
      <div class="button-group">
        <button class="btn add-btn" @click="goToAdd">➕ Add Employee</button>
        <button class="btn list-btn" @click="goToList">📋 Employee List</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router"

const router = useRouter()

const goToAdd = () => {
  router.push("/add")
}

const goToList = () => {
  router.push("/list")
}
</script>

<style scoped>
/* Full screen background */
.home-container {
  height: 100vh;
  width: 100%;
  background-image: url("https://images.unsplash.com/photo-1521737604893-d14cc237f11d");
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Dark overlay for readability */
.overlay {
  background: rgba(0, 0, 0, 0.5);
  padding: 60px;
  border-radius: 16px;
  text-align: center;
}

/* Centered buttons */
.button-group {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* Button base style */
.btn {
  padding: 16px 40px;
  border: none;
  border-radius: 10px;
  font-size: 18px;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.3);
}

/* Add Employee button */
.add-btn {
  background: linear-gradient(135deg, #007bff, #0056b3);
  color: white;
}
.add-btn:hover {
  background: linear-gradient(135deg, #3399ff, #004085);
  transform: scale(1.05);
}

/* Employee List button */
.list-btn {
  background: linear-gradient(135deg, #28a745, #1c7c31);
  color: white;
}
.list-btn:hover {
  background: linear-gradient(135deg, #4cd964, #146b24);
  transform: scale(1.05);
}
</style>
