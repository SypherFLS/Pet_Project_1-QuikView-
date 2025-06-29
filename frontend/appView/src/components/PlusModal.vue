<template>
  <div v-if="isOpen" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>Добавить новый элемент</h2>
        <button class="close-btn" @click="closeModal">×</button>
      </div>
      
      <div class="modal-body">
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
        
        <div class="form-group">
          <label for="description">Описание</label>
          <textarea 
            id="description" 
            v-model="formData.description" 
            placeholder="Введите описание"
            class="form-textarea"
            rows="3"
          ></textarea>
        </div>
        
        <div class="form-group">
          <label for="category">Категория</label>
          <select id="category" v-model="formData.category" class="form-select">
            <option value="">Выберите категорию</option>
            <option value="work">Работа</option>
            <option value="personal">Личное</option>
            <option value="project">Проект</option>
            <option value="other">Другое</option>
          </select>
        </div>
      </div>
      
      <div class="modal-footer">
        <button class="btn btn-secondary" @click="closeModal">Отмена</button>
        <button class="btn btn-primary" @click="handleSubmit">Добавить</button>
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

const formData = reactive({
  title: '',
  description: '',
  category: ''
})

const closeModal = () => {
  emit('close')
  // Сброс формы
  formData.title = ''
  formData.description = ''
  formData.category = ''
}

const handleSubmit = () => {
  if (formData.title.trim()) {
    emit('submit', { ...formData })
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
  border: 1px solid rgba(138, 43, 226, 0.4);
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  max-height: 80vh;
  overflow-y: auto;
  box-shadow: 
    0 10px 30px rgba(0, 0, 0, 0.5),
    0 0 30px rgba(138, 43, 226, 0.3);
  animation: slideIn 0.3s ease;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 25px;
  border-bottom: 1px solid rgba(138, 43, 226, 0.3);
}

.modal-header h2 {
  color: #e0e0e0;
  margin: 0;
  font-size: 1.5rem;
  text-shadow: 0 0 10px rgba(138, 43, 226, 0.5);
}

.close-btn {
  background: none;
  border: none;
  color: #e0e0e0;
  font-size: 24px;
  cursor: pointer;
  padding: 5px;
  border-radius: 50%;
  width: 35px;
  height: 35px;
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
  padding: 25px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  color: #e0e0e0;
  margin-bottom: 8px;
  font-weight: 500;
  text-shadow: 0 0 5px rgba(138, 43, 226, 0.3);
}

.form-input,
.form-textarea,
.form-select {
  width: 100%;
  padding: 12px 15px;
  background: linear-gradient(145deg, #2a1a3e, #1a0a2e);
  border: 1px solid rgba(138, 43, 226, 0.4);
  border-radius: 8px;
  color: #e0e0e0;
  font-size: 14px;
  transition: all 0.2s ease;
  box-sizing: border-box;
}

.form-input:focus,
.form-textarea:focus,
.form-select:focus {
  outline: none;
  border-color: #8a2be2;
  box-shadow: 
    0 0 0 2px rgba(138, 43, 226, 0.2),
    0 0 15px rgba(138, 43, 226, 0.3);
}

.form-input::placeholder,
.form-textarea::placeholder {
  color: #888;
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 20px 25px;
  border-top: 1px solid rgba(138, 43, 226, 0.3);
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
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
  box-shadow: 
    0 0 20px rgba(138, 43, 226, 0.4),
    0 0 30px rgba(138, 43, 226, 0.2);
  transform: translateY(-1px);
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

/* Адаптивность */
@media (max-width: 768px) {
  .modal-content {
    width: 95%;
    margin: 10px;
  }
  
  .modal-header,
  .modal-body,
  .modal-footer {
    padding: 15px;
  }
}
</style> 