import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ShoppingCart } from 'lucide-react'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
})

function HomePage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <nav className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-2">
              <ShoppingCart className="w-8 h-8 text-blue-600" />
              <span className="text-2xl font-bold text-gray-900">Gophiway</span>
            </div>
            <div className="flex space-x-4">
              <button className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">
                Login
              </button>
              <button className="bg-blue-600 text-white hover:bg-blue-700 px-4 py-2 rounded-md text-sm font-medium">
                Sign Up
              </button>
            </div>
          </div>
        </div>
      </nav>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-gray-900 mb-4">
            Welcome to <span className="text-blue-600">Gophiway</span>
          </h1>
          <p className="text-xl text-gray-600 mb-8">The fast lane for modern e-commerce</p>
          <div className="bg-white rounded-lg shadow-lg p-8 max-w-2xl mx-auto">
            <h2 className="text-2xl font-semibold text-gray-800 mb-4">🚀 Framework Ready!</h2>
            <div className="text-left space-y-3 text-gray-600">
              <p>✅ Go backend with Fiber framework</p>
              <p>✅ React frontend with TypeScript</p>
              <p>✅ Tailwind CSS 4 styling</p>
              <p>✅ PostgreSQL database</p>
              <p>✅ Redis caching</p>
              <p>✅ Docker development environment</p>
              <p>✅ Security-first architecture</p>
            </div>
            <div className="mt-6 pt-6 border-t border-gray-200">
              <p className="text-sm text-gray-500">
                Backend running on:{' '}
                <code className="bg-gray-100 px-2 py-1 rounded">http://localhost:8080</code>
              </p>
              <p className="text-sm text-gray-500 mt-2">
                Frontend running on:{' '}
                <code className="bg-gray-100 px-2 py-1 rounded">http://localhost:5173</code>
              </p>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  )
}

export default App
