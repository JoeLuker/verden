<script>
	function handleClick() {
		alert('no more alerts');
	}
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>


	<h1>
		Hi Soph!
	</h1>

	<button on:click={handleClick}> Click me </button>

</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	button {
    background-color: #ff3e00; /* Svelte's signature orange */
    border: 2px solid transparent; /* Transparent border */
    color: white; /* White text */
    padding: 10px 20px; /* Comfortable padding */
    text-align: center; /* Centered text */
    text-decoration: none; /* No underline */
    display: inline-block; /* Inline-block display */
    font-size: 16px; /* Suitable font size */
    font-weight: bold; /* Bold text */
    margin: 5px; /* Slight margin */
    cursor: pointer; /* Pointer cursor on hover */
    border-radius: 4px; /* Slightly rounded corners */
    transition: background-color 0.3s, border-color 0.3s; /* Smooth transition for background and border color */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); /* Subtle shadow for depth */
	}

	button:hover, button:focus {
		background-color: #e63900; /* Slightly darker orange when hovered/focused */
		border-color: #fff; /* White border on hover/focus */
		outline: none; /* Remove outline */
	}

</style>
