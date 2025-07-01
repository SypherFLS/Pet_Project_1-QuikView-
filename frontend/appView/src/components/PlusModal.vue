<template>
  <div v-if="isOpen" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <div class="tabs">
          <button
            :class="['tab-btn', { active: selectedTab === 'interview' }]"
            @click="selectedTab = 'interview'"
          >Интервью</button>
          <button
            :class="['tab-btn', { active: selectedTab === 'mockup' }]"
            @click="selectedTab = 'mockup'"
          >Макет</button>
        </div>
        <button class="close-btn" @click="closeModal">×</button>
      </div>
      
      <div class="modal-body">
        <template v-if="selectedTab === 'interview'">
          <div class="form-group">
            <label for="interviewee">Имя интервьюируемого</label>
            <input
              type="text"
              id="interviewee"
              v-model="formData.interviewee"
              placeholder="Введите имя"
              class="form-input"
            >
          </div>
        </template>
        <template v-else>
          <div class="form-group">
            <label for="title">Название</label>
            <input
              type="text"
              id="title"
              v-model="formData.title"
              placeholder="Введите название"
              class="form-input"
            >
          </div>
        </template>
      </div>
      
      <div class="modal-footer">
        <button class="btn btn-secondary" @click="closeModal">Отмена</button>
        <button class="btn btn-primary" @click="handleSubmit">Создать</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const props = defineProps({
  isOpen: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'submit'])

const selectedTab = ref('interview')

const formData = reactive({
  interviewee: '',
  title: ''
})

const closeModal = () => {
  emit('close')
  // Сброс формы
  formData.interviewee = ''
  formData.title = ''
  selectedTab.value = 'interview'
}

const handleSubmit = () => {
  if (
    (selectedTab.value === 'interview' && formData.interviewee.trim()) ||
    (selectedTab.value === 'mockup' && formData.title.trim())
  ) {
    emit('submit', {
      type: selectedTab.value,
      interviewee: formData.interviewee,
      title: formData.title
    })
    closeModal()
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  animation: fadeIn 0.2s ease;
}

.modal-content {
  background: linear-gradient(145deg, #1a0a2e, #0a0a0a);
  border: 1px solid rgba(138, 43, 226, 0.32);
  border-radius: 20px;
  width: 100%;
  max-width: 420px;
  min-width: 320px;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 0 12px 36px rgba(0,0,0,0.55), 0 0 32px rgba(138,43,226,0.18);
  animation: slideIn 0.3s ease;
  padding: 32px 32px 24px 32px;
  margin-top: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.modal-header {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 0 0 12px 0;
  border-bottom: 1px solid rgba(138, 43, 226, 0.10);
}

.tabs {
  display: flex;
  gap: 16px;
  justify-content: center;
  align-items: center;
  width: 100%;
  margin: 0 0 18px 0;
  border-bottom: 1.5px solid rgba(138, 43, 226, 0.10);
}
.tab-btn {
  background: linear-gradient(145deg, #2a1a3e, #1a0a2e);
  color: #e0e0e0;
  border: none;
  border-radius: 12px 12px 0 0;
  padding: 12px 32px;
  font-size: 1.12rem;
  font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.18s cubic-bezier(.4,2,.6,1);
  outline: none;
  margin-bottom: -2px;
  box-shadow: none;
  letter-spacing: 0.1px;
  opacity: 0.95;
  position: relative;
  line-height: 1.3;
}
.tab-btn.active {
  background: linear-gradient(145deg, #8a2be2 80%, #6a1b9a 100%);
  color: #fff;
  border-bottom: 3px solid #8a2be2;
  z-index: 1;
  opacity: 1;
  box-shadow: 0 3px 12px rgba(138,43,226,0.13);
  font-weight: 600;
}

.close-btn {
  background: none;
  border: none;
  color: #e0e0e0;
  font-size: 28px;
  cursor: pointer;
  padding: 5px;
  border-radius: 50%;
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}
.close-btn:hover {
  background: rgba(138, 43, 226, 0.2);
  color: #ffffff;
  text-shadow: 0 0 10px rgba(138, 43, 226, 0.8);
}

.modal-body {
  width: 100%;
  padding: 28px 0 0 0;
  display: flex;
  flex-direction: column;
  gap: 24px;
  align-items: center;
}

.form-group {
  width: 100%;
  margin-bottom: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-size: 1.08rem;
  font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
  line-height: 1.4;
  color: #e0e0e0;
  font-weight: 500;
  text-shadow: 0 0 5px rgba(138, 43, 226, 0.13);
}

.form-input {
  font-size: 1.08rem;
  font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
  line-height: 1.4;
  padding: 13px 18px;
  border-radius: 9px;
  border: 1.2px solid rgba(138, 43, 226, 0.22);
  background: linear-gradient(145deg, #2a1a3e, #1a0a2e);
  color: #e0e0e0;
  width: 100%;
  transition: all 0.2s ease;
  box-sizing: border-box;
}
.form-input:focus {
  outline: none;
  border-color: #8a2be2;
  box-shadow: 0 0 0 2px rgba(138, 43, 226, 0.13), 0 0 15px rgba(138, 43, 226, 0.13);
}
.form-input::placeholder {
  color: #888;
}

.modal-footer {
  width: 100%;
  display: flex;
  justify-content: center;
  gap: 0;
  padding: 28px 0 0 0;
  border-top: none;
}

.btn {
  width: 100%;
  padding: 15px 0;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 1.13rem;
  font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
  font-weight: 600;
  transition: all 0.2s ease;
}

.btn-secondary {
  background: linear-gradient(145deg, #3a2a4e, #2a1a3e);
  color: #e0e0e0;
  border: 1px solid rgba(138, 43, 226, 0.4);
}

.btn-secondary:hover {
  background: linear-gradient(145deg, #4a3a5e, #3a2a4e);
  border-color: rgba(138, 43, 226, 0.6);
  box-shadow: 0 0 15px rgba(138, 43, 226, 0.2);
}

.btn-primary {
  background: linear-gradient(145deg, #8a2be2, #6a1b9a);
  color: #ffffff;
  border: 1px solid #8a2be2;
}
.btn-primary:hover {
  background: linear-gradient(145deg, #9a3bf2, #7a2baa);
  box-shadow: 0 0 20px rgba(138, 43, 226, 0.18), 0 0 30px rgba(138, 43, 226, 0.09);
  transform: translateY(-1px) scale(1.03);
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideIn {
  from { 
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to { 
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@media (max-width: 768px) {
  .modal-content {
    width: 98vw;
    min-width: unset;
    max-width: 99vw;
    padding: 10px 4vw 10px 4vw;
    margin-top: 8px;
  }
  .modal-header,
  .modal-body,
  .modal-footer {
    padding-left: 0;
    padding-right: 0;
  }
  .tabs {
    gap: 8px;
    margin: 0 0 10px 0;
  }
  .tab-btn {
    padding: 10px 10vw;
    font-size: 1.05rem;
  }
  .modal-body {
    padding: 10px 0 0 0;
    gap: 14px;
  }
  .btn {
    padding: 13px 0;
    font-size: 1.08rem;
  }
}

.add-position {
  display: flex;
  gap: 8px;
  margin-top: 8px;
}
</style> 