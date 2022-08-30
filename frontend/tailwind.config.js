/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                'blurple': '#5865F2',
            }
        },
    },
    plugins: [
        require('@tailwindcss/forms')
    ],
}
