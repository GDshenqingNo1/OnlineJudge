<template>
    <div>
      <ProblemList></ProblemList>
    </div>
</template>


<script>
import Navi from "./Navi.vue";
import ProblemList from "./ProblemList.vue";
export default  {
  name:"Home",
  components: {ProblemList, Navi},
  data () {
  }
}
</script>

