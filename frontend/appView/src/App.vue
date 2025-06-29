<template>
  <div>
    <navBar @openPlusModal="showPlusModal = true" />
    <router-view />
    <PlusModal 
      :isOpen="showPlusModal" 
      @close="showPlusModal = false"
      @submit="handlePlusSubmit"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import navBar from './components/navBar.vue'
import PlusModal from './components/PlusModal.vue'

const showPlusModal = ref(false)

const handlePlusSubmit = (data) => {
  console.log('Новый элемент:', data)
  // Здесь можно добавить логику сохранения данных
}
</script>

<style scoped>

</style>
