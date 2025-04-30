<template>
  <div class="message-box" :style="{ borderColor: borderColor }">
    <h2>{{ title }}</h2>
    <p>{{ message }}</p>
    <input v-model="newMessage" placeholder="输入新留言" />
    <button @click="updateMessage">更新留言</button>
  </div>
</template>

<script>
export default {
  name: 'MessageBox',
  props: {
    title: {
      type: String,
      default: '留言板'
    },
    borderColor: {
      type: String,
      default: '#42b983'  // Vue 绿色
    }
  },
  data() {
    return {
      message: '你好，这是默认留言',
      newMessage: ''
    }
  },
  methods: {
    updateMessage() {
      if (this.newMessage.trim()) {
        this.message = this.newMessage
        this.newMessage = ''
      }
    }
  }
}
</script>

<style>
.message-box {
  border: 2px solid;
  padding: 12px;
  margin: 16px 0;
  border-radius: 8px;
}
</style>