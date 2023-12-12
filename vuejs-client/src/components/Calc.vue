<template>
	<div class="calc-wrapper">
	  <h1 class="item">Pack Calculation</h1>
	  <div class="item">
		<input class="input" type="number" pattern="/^\d+$/" min="0" v-model.number="inputData" placeholder="Enter items" />
		<button class="button" :disabled="isDisabled" @click="makeRequest">Calculate</button>
	  </div>
	  <div v-if="response" class="item">
		<div v-for="pack in response">
			<div>Size: {{pack.size}} - Quantity: {{pack.quantity}}</div>
		</div>
	  </div>
	</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      inputData: '',
      response: null,
    };
  },
  computed: {
	isDisabled() {
		return this.inputData == '' || this.inputData <= 0 || this.inputData % 1 != 0;
	}
  },
  methods: {
    async makeRequest() {
		
		if (this.inputData == '') {
			return
		}

		try {
			const apiUrl = 'http://3.15.189.102:3000/calculate?items='+this.inputData;
			const response = await axios.get(apiUrl);
			this.response = response.data;
		} catch (error) {
			console.error('Error fetching data:', error);
			this.response = null;
		}
    },
  },
};
</script>

<style scoped>
	.calc-wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.input {
		border-radius: 8px;
		border-width: 0;
		display: inline-block;
		font-size: 14px;
		font-weight: 500;
		line-height: 20px;
		margin: 10px;
		padding: 10px 12px 10px 10px;
	}

	.button {
		border-radius: 8px;
		border-width: 0;
		cursor: pointer;
		display: inline-block;
		font-size: 14px;
		font-weight: 500;
		line-height: 20px;
		list-style: none;
		margin: 10px 5px;
		padding: 10px 12px 10px 10px;
		border: 1px solid green;
		text-align: center;
	}

</style>