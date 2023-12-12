<script setup>
import Calc from './components/Calc.vue'
import Pack from './components/Pack.vue'
</script>

<template>
  <div class="main-wrapper">
    <Pack class="item"/>
    <Calc class="item"/>
    
  </div>
</template>

<style scoped>

	@media only screen and (max-width: 768px) {
		/* For mobile phones: */
		.main-wrapper {	
			flex-direction: column !important;
		}
	}


   .main-wrapper {	
		display: flex;
		flex-direction: row;
		align-items: center;
		align-content: space-between;
		vertical-align: center;
	}

	.item {
		width: 300px;
		margin: 30px;
	}

</style>
