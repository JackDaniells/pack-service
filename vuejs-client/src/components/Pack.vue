<template>
	<div class="pack-wrapper">
		<h1 class="item">Pack List</h1>
		<div class="item">
			<div class="row">
				<input class="input" type="number" pattern="/^\d+$/" min="0" v-model.number="inputData" placeholder="Enter pack size" />
				<button class="button" :disabled="isDisabled" @click="addPack">Add</button>
				<button class="button" :disabled="isDisabled" @click="removePack">Remove</button>
			</div>
		</div>
		<div v-if="packs" class="item">
			<ul v-for="pack in packs">
				<li>{{pack.size}}</li>
			</ul>
		</div>
	</div>
</template>

<script>
import axios from 'axios';

export default {
  mounted() {
    this.getAllPacks()
  },
  data() {
    return {
      inputData: '',
      packs: null,
    };
  },
  computed: {
	isDisabled() {
		return this.inputData == '' || this.inputData <= 0 || this.inputData % 1 != 0;
	}
  },
  methods: {
	async getAllPacks() {
		try {
			const apiUrl = 'http://3.15.189.102:3000/packs';
			const response = await axios.get(apiUrl);
			this.packs = response.data;
		} catch (error) {
			console.error('Error fetching data:', error);
		}
    },

    async addPack() {
		try {
			const apiUrl = 'http://3.15.189.102:3000/packs';
			await axios.post(apiUrl, { "size": parseInt(this.inputData) });
			await this.getAllPacks()
		} catch (error) {
			console.error('Error fetching data:', error);
		}
    },

	async removePack() {
		try {
			const apiUrl = 'http://3.15.189.102:3000/packs/' + this.inputData;
			await axios.delete(apiUrl);
			await this.getAllPacks()
		} catch (error) {
			console.error('Error fetching data:', error);
		}
    },
  },
};
</script>

<style scoped>

	.pack-wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.row {
		display: flex;
		flex-direction: row;
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
		margin: 10px 0px;
		padding: 10px 12px 10px 10px;
		text-align: center;
		border: 1px solid green;
	}
</style>