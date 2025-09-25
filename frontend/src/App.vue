<template>
  <div id="app" class="min-h-screen bg-gray-100">
    <header class="bg-blue-600 text-white py-4 shadow-md">
      <center>
      <h1 class="text-center text-2xl font-bold">Employee Management System</h1>
      </center>
    </header>

    <main class="p-6">
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: "App",
}
</script>

<style>
body {
  margin: 0;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
}
</style>
