<script setup>
defineProps({
  categories: {
    type: Array,
    required: true,
  },
  activeCategory: {
    type: String,
    required: true,
  },
  keyword: {
    type: String,
    required: true,
  },
})

const emit = defineEmits(['change-category', 'update-keyword'])
</script>

<template>
  <div class="message-toolbar" aria-label="资讯筛选">
    <div class="message-tabs">
      <button
        v-for="category in categories"
        :key="category"
        class="message-tab"
        :class="{ active: category === activeCategory }"
        type="button"
        @click="emit('change-category', category)"
      >
        {{ category }}
      </button>
    </div>

    <label class="message-search">
      <span aria-hidden="true"></span>
      <input
        :value="keyword"
        type="search"
        placeholder="请输入关键词搜索资讯"
        aria-label="请输入关键词搜索资讯"
        @input="emit('update-keyword', $event.target.value)"
      />
      <i aria-hidden="true"></i>
    </label>
  </div>
</template>

<style scoped>
.message-toolbar {
  width: 1200px;
  height: 96px;
  margin: 0 auto;
  padding-top: 20px;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
}

.message-tabs {
  display: flex;
  align-items: flex-start;
}

.message-tab {
  height: 56px;
  border: 0;
  padding: 0 16px;
  background: transparent;
  color: #333;
  font-size: 16px;
  line-height: 56px;
  font-weight: 400;
}

.message-tab:first-child {
  padding-left: 0;
}

.message-tab.active,
.message-tab:hover {
  color: #ff4848;
}

.message-tab.active {
  font-weight: 600;
}

.message-search {
  width: 190px;
  height: 34px;
  margin-top: 11px;
  display: flex;
  align-items: center;
  color: #333;
}

.message-search span,
.message-search i {
  position: relative;
  width: 16px;
  height: 16px;
  flex: 0 0 auto;
}

.message-search span::before {
  content: "";
  position: absolute;
  left: 1px;
  top: 1px;
  width: 10px;
  height: 10px;
  border: 2px solid currentColor;
  border-radius: 50%;
}

.message-search span::after {
  content: "";
  position: absolute;
  right: 1px;
  bottom: 1px;
  width: 7px;
  height: 2px;
  background: currentColor;
  transform: rotate(45deg);
  transform-origin: center;
}

.message-search i::before,
.message-search i::after {
  content: "";
  position: absolute;
  left: 3px;
  top: 7px;
  width: 10px;
  height: 1px;
  background: currentColor;
}

.message-search i::before {
  transform: rotate(45deg);
}

.message-search i::after {
  transform: rotate(-45deg);
}

.message-search input {
  width: 144px;
  height: 32px;
  border: 0;
  padding: 0;
  outline: 0;
  color: #333;
  background: transparent;
  font-size: 14px;
  line-height: 32px;
}

.message-search input::placeholder {
  color: #333;
}

@media (max-width: 900px) {
  .message-toolbar {
    width: calc(100% - 32px);
    height: auto;
    padding: 20px 0 24px;
    flex-direction: column;
    gap: 14px;
  }

  .message-tabs {
    width: 100%;
    overflow-x: auto;
  }

  .message-tab {
    flex: 0 0 auto;
  }

  .message-search {
    width: 100%;
    margin-top: 0;
    padding: 0 12px;
    border: 1px solid #ececec;
    border-radius: 4px;
  }

  .message-search input {
    flex: 1;
    width: auto;
  }
}
</style>
