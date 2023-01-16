import './App.css';
import "bootstrap/dist/css/bootstrap.min.css";
import MySideNav from './components/MySideNav';
import TopNav from './components/TopNav';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import Deployment from './pages/Deployment';
import Troubleticket from './pages/Troubleticket';
import Reports from './pages/Reports';

function App() {
  return (
    <>
    <MySideNav>
    <TopNav />
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Dashboard />} />
        <Route path="/deployment" element={<Deployment />} />
        <Route path="/troubletickets" element={<Troubleticket />} />
        <Route path="/reports" element={<Reports />} />
        </Routes>
    </BrowserRouter>
    </MySideNav>
    </>
  );
}

export default App;
