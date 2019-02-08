<template lang="pug">
nav(class="navbar navbar-light")
  div(class="container")
    router-link(class="navbar-brand" :to="{ name: 'home' }") example
    ul(class="nav navbar-nav pull-xs-right")
      li(class="nav-item")
        router-link(
          class="nav-link"
          active-class="active"
          exact :to="{ name: 'home' }") Home
</template>

<script>
export default {
  name: 'main-header',
  computed: {}
};
</script>
