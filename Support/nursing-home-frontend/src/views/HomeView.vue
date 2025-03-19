<template>
  <div class="home">
    <div class="hero">
      <h1>找到最适合的养老院</h1>
      <p>我们帮助您找到最合适的养老机构</p>
      <router-link to="/nursing-homes" class="cta-button">
        开始查找
      </router-link>
    </div>

    <div class="features">
      <div class="feature">
        <div class="icon">🏠</div>
        <h3>优质机构</h3>
        <p>严格筛选，确保服务质量</p>
      </div>
      <div class="feature">
        <div class="icon">💰</div>
        <h3>透明价格</h3>
        <p>清晰的价格信息展示</p>
      </div>
      <div class="feature">
        <div class="icon">👥</div>
        <h3>专业服务</h3>
        <p>专业的养老护理团队</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  min-height: calc(100vh - 64px);
}

.hero {
  text-align: center;
  padding: 80px 20px;
  background: linear-gradient(to bottom, #e8f5e9, #fff);
}

.hero h1 {
  font-size: 2.5em;
  margin-bottom: 20px;
  color: #2c3e50;
}

.hero p {
  font-size: 1.2em;
  color: #666;
  margin-bottom: 30px;
}

.cta-button {
  display: inline-block;
  padding: 15px 30px;
  background: #4caf50;
  color: white;
  text-decoration: none;
  border-radius: 30px;
  font-weight: bold;
  transition: background 0.3s;
}

.cta-button:hover {
  background: #388e3c;
}

.features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  padding: 60px 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.feature {
  text-align: center;
  padding: 30px;
}

.icon {
  font-size: 3em;
  margin-bottom: 20px;
}

.feature h3 {
  margin-bottom: 15px;
  color: #2c3e50;
}

.feature p {
  color: #666;
}
</style> 