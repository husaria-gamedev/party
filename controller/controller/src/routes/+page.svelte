<main class="grid absolute h-screen w-full place-items-center bg-gray-900">
    <form action="" method="post" class="flex flex-col gap-8 m-auto items-center rounded-md">
        <div>
            <label for="code" class="self-start p-4 my-8 text-2xl text-white">Code:</label><br/>
            <input type="text" name="code" id="code" placeholder="Room code" class="bg-white rounded p-3 border-none bg-transparent text-lg text-gray-900 focus:outline-none my-4">
        </div>
        <div>
            <label for="name" class="self-start p-4 text-2xl text-white">Name:</label><br/>
            <input type="text" name="name" id="name" placeholder="Name that other players will see" class="bg-white rounded p-3 border-none bg-transparent text-lg text-gray-900 focus:outline-none my-4">
        </div>
        <button type="submit" class="relative inline-flex items-center justify-center p-4 px-6 py-3 overflow-hidden font-medium text-indigo-600 transition duration-300 ease-out border-2 border-orange rounded-full shadow-md group">
            <span class="absolute inset-0 flex items-center justify-center w-full h-full text-white duration-300 -translate-x-full bg-orange group-hover:translate-x-0 ease">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path></svg>
            </span>
            <span class="absolute flex items-center justify-center w-full h-full text-orange transition-all duration-300 transform group-hover:translate-x-full ease">PLAY</span>
            <span class="relative invisible">PLAY NOW</span>
        </button>
    </form>
</main>
<header class="absolute flex flex-row justify-center gap-2 bg-white w-full p-2 text-center z-2">
    <!-- game logo or and name -->
    <div>PARTY</div>
    <!-- help icon -->
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 5.25h.008v.008H12v-.008Z" />
      </svg>
</header>