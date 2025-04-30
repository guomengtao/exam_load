
<template>
  <div id="app">
   <h1>我的 Vue 应用</h1>
    <nav>
      <router-link to="/">首页</router-link> |
      <router-link to="/about">关于我们</router-link>
    </nav>
    <router-view />
    <MessageBox title="留言区 A" borderColor="#42b983" />
    <MessageBox title="留言区 B" borderColor="#e74c3c" />
  </div>
</template>

<script>
import MessageBox from './components/MessageBox.vue'

export default {
  name: 'App',
  components: {
    MessageBox
  }
}
</script>