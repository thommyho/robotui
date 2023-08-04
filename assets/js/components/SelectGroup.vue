<template>
	<div class="mode-group border d-inline-flex" role="group">
		<button
			v-for="(option, i) in options"
			:id="i === 0 ? id : null"
			:key="option.value"
			type="button"
			class="btn btn-sm flex-grow-1 flex-shrink-1"
			:class="{ active: option.value === modelValue }"
			:disabled="option.disabled"
			@click="$emit('update:modelValue', option.value)"
		>
			{{ option.name }}
		</button>
	</div>
</template>

<script>
export default {
	name: "SelectGroup",
	props: {
		id: String,
		options: Array,
		modelValue: [Number, String],
	},
	emits: ["update:modelValue"],
};
</script>

<style scoped>
.mode-group {
	border: 2px solid var(--robotui-default-text);
	border-radius: 17px;
	padding: 4px;
}

.btn {
	/* equal width buttons */
	flex-basis: 0;
	white-space: nowrap;
	border-radius: 12px;
	padding: 0.1em 0.8em;
	color: var(--robotui-default-text);
	border: none;
	overflow-x: hidden;
	text-overflow: ellipsis;
}
.btn:hover {
	color: var(--robotui-gray);
}
.btn.active {
	color: var(--robotui-background);
	background: var(--robotui-default-text);
}
.btn-group {
	border-radius: 16px;
}
</style>
