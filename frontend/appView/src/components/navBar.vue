<template>
  <nav class="vertical-navbar">
    <div class="nav-section top-section">
      <button class="nav-item circle-item" @click="openPlusModal">
        <span class="nav-icon">➕</span>
      </button>
      <router-link to="/profiles" class="nav-item circle-item" active-class="active-link">
        <span class="nav-icon">👥</span>
      </router-link>
      <router-link to="/stats" class="nav-item circle-item" active-class="active-link">
        <span class="nav-icon">📊</span>
      </router-link>
    </div>
    
    <div class="nav-section bottom-section">
      <router-link to="/settings" class="nav-item circle-item" active-class="active-link">
        <span class="nav-icon">⚙️</span>
      </router-link>
      <router-link to="/account" class="nav-item circle-item" active-class="active-link">
        <span class="nav-icon">👤</span>
      </router-link>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['openPlusModal'])

const openPlusModal = () => {
  emit('openPlusModal')
}
</script>

<style scoped>
.vertical-navbar {
  position: fixed;
  right: 0;
  top: 0;
  height: 100vh;
  width: 70px;
  background: linear-gradient(180deg, #0a0a0a 0%, #1a0a2e 50%, #0a0a0a 100%);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 20px 0;
  box-shadow: -2px 0 15px rgba(138, 43, 226, 0.2);
  z-index: 1000;
  border-left: 1px solid rgba(138, 43, 226, 0.3);
}

.nav-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
}

.top-section {
  margin-top: 20px;
}

.bottom-section {
  margin-bottom: 20px;
}

.nav-item {
  cursor: pointer;
  transition: all 0.2s ease;
  text-decoration: none;
  display: block;
}

.circle-item {
  width: 45px;
  height: 45px;
  border-radius: 50%;
  background: linear-gradient(145deg, #1a0a2e, #0a0a0a);
  border: 1px solid rgba(138, 43, 226, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  box-shadow: 
    inset 0 1px 0 rgba(255, 255, 255, 0.1),
    0 2px 8px rgba(0, 0, 0, 0.3);
}

.circle-item:hover {
  background: linear-gradient(145deg, #2a1a3e, #1a0a2e);
  border-color: rgba(138, 43, 226, 0.7);
  transform: scale(1.05);
  box-shadow: 
    inset 0 1px 0 rgba(255, 255, 255, 0.15),
    0 4px 12px rgba(138, 43, 226, 0.3),
    0 0 20px rgba(138, 43, 226, 0.2);
}

.active-link {
  background: linear-gradient(145deg, #3a2a4e, #2a1a3e) !important;
  border-color: #8a2be2;
  transform: scale(1.08);
  box-shadow: 
    inset 0 1px 0 rgba(255, 255, 255, 0.2),
    0 6px 16px rgba(138, 43, 226, 0.4),
    0 0 25px rgba(138, 43, 226, 0.3);
}

.nav-icon {
  font-size: 18px;
  color: #e0e0e0;
  text-shadow: 0 0 8px rgba(138, 43, 226, 0.5);
  transition: all 0.2s ease;
}

.circle-item:hover .nav-icon {
  color: #f0f0f0;
  text-shadow: 0 0 12px rgba(138, 43, 226, 0.8);
}

.active-link .nav-icon {
  color: #ffffff;
  text-shadow: 0 0 15px rgba(138, 43, 226, 1);
}

/* Адаптивность для мобильных устройств */
@media (max-width: 768px) {
  .vertical-navbar {
    width: 60px;
  }
  
  .circle-item {
    width: 40px;
    height: 40px;
  }
  
  .nav-icon {
    font-size: 16px;
  }
}
</style>