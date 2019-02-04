<template>
  <div id="app">
    <MainHeader />
    <router-view />
    <MainFooter />
  </div>
</template>

<script>
import MainHeader from '@/components/main-header';
import MainFooter from '@/components/main-footer';

export default {
  name: 'app',
  components: {MainFooter, MainHeader}
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
