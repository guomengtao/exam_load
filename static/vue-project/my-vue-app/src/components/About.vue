<!-- src/views/About.vue -->
<template>
  <div>
    <h2>这是 About 页面</h2>
    <MessageBox msg="这是从 About 页面传入的 message" />
  </div>
</template>

<script>
import MessageBox from '@/components/MessageBox.vue'

export default {
  name: 'About',
  components: {
    MessageBox
  }
}
</script>