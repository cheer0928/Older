<template>
  <nav class="nav-bar">
    <router-link to="/" class="logo">
      养老院导航
    </router-link>
    <div class="nav-links">
      <router-link to="/nursing-homes">养老院列表</router-link>
      <router-link to="/about">关于我们</router-link>
    </div>
  </nav>
</template>

<style scoped>
.nav-bar {
  background: #fff;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
  text-decoration: none;
}

.nav-links {
  display: flex;
  gap: 20px;
}

.nav-links a {
  color: #666;
  text-decoration: none;
  padding: 5px 10px;
}

.nav-links a:hover {
  color: #4CAF50;
}
</style> 