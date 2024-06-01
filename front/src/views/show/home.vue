<template>
  <el-container class="top-container">
    <el-main class="main-content">
      <h1 class="headline">区块链非物质遗产平台</h1>
      <div class="entry-grid">
        <div class="entry-box" @click="$router.push('/look')">
          <h2>非遗概览</h2>
          <!-- You can add an icon here -->
        </div>
        <div class="entry-box" @click="$router.push('/edit')">
          <h2>申请非遗</h2>
          <!-- Icon if needed -->
        </div>
        <div class="entry-box" @click="$router.push('/admin')">
          <h2>登录入口</h2>
          <!-- Login icon -->
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'HomePage'
}
</script>

<style scoped>
.top-container {
  height: 100vh;
  overflow: hidden;
  background-image: linear-gradient(to right bottom, #8e44ad, #f1c40f); /* 更改为较浅的紫色到黄色渐变 */
}

.main-content {
  display: flex;
  margin-top: 100px;
  align-items: center;
  height: 100%;
  text-align: center;
  color: white;
  flex-direction: column;
}

.entry-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 30px;
}

.entry-box {
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  padding: 30px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

.entry-box:hover {
  background-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-5px);
}

h2 {
  margin-bottom: 10px;
  font-weight: 300;
}

.headline {
  font-family: 'Arial', sans-serif; /* 可替换为您喜欢的字体 */
  font-size: 48px; /* 根据需要调整字体大小 */
  font-weight: bold;
  color: #ffffff; /* 文字颜色，这里假设背景是深色所以用了白色 */
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5); /* 添加阴影效果，增强立体感 */
  text-align: center; /* 居中显示 */
  margin-top: 100px; /* 根据页面布局调整顶部外边距 */
  letter-spacing: 2px; /* 字符间距，使文字更加优雅 */
  /* 下面是可选的渐变效果，为文字添加色彩变化 */
  background-color: white;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  text-align: center;
}
</style>