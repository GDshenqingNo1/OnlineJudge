<template>
  <div id="app">
    <router-view :key="$route.fullPath"></router-view>
  </div>
</template>

<script>

export default {
  watch: {
    $route() {
      location.reload();
    },
  },
  name: 'App'
}
</script>

