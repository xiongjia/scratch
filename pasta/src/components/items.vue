<template lang="pug">
div
  ul(class="item-data")
    li(v-for="(value, key) in testItems" :key="key")
      span(v-text="key")
      span(v-for="data in value" :key="data" v-text="data")
</template>

<script>
export default {
  name: 'items',
  props: {
    testItems: { type: Object, required: true }
  }
};
</script>
