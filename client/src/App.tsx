import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './pages/Login';
import Home from './pages/Home';
import Register from './pages/Register';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path='/authen' element={<Login/>}/>
        <Route path='/' element={<Home/>}/>
        <Route path='/create-user' element={<Register/>}/>
      </Routes>
    </Router>
  );
};

export default App;
 