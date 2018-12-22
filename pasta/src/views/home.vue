<template>
  <div class='home'>
    <!-- banner -->
    <div class="banner">
      <div class="container">
        <h1 class="logo-font">title</h1>
        <p>description</p>
      </div>
    </div>

    <!-- page -->
    <div class="container page">
      <keep-alive>
        <component ref="component" :is="currentView" />
      </keep-alive>
    </div>
  </div>
</template>

<script>
import Info from '@/views/info.vue';
import Detail from '@/views/detail.vue';

export default {
  name: 'home',
  components: { Info, Detail },
  data () {
    return { currentView: Info };
  },
  created () {
    this.$log.debug('view(home) is created');
  }
};
</script>
