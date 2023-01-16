import React from 'react'
import SideNav, {Toggle, NavItem, NavIcon, NavText} from '@trendmicro/react-sidenav'
import {Link} from 'react-router-dom';
import {MdOutlineDashboard} from 'react-icons/md';
import {GiServerRack} from 'react-icons/gi';
import {FaBug, FaProjectDiagram} from 'react-icons/fa';
import {TbSubtask} from 'react-icons/tb';
import {HiOutlineDocumentReport} from 'react-icons/hi'
import "@trendmicro/react-sidenav/dist/react-sidenav.css"

const MySideNav = () => {
  return (
    <SideNav
        onSelect={selected =>{
            console.log(selected)
        }}>
            <SideNav.Toggle />
            <SideNav.Nav defaultSelected="home">
                <NavItem eventKey="dashboard">
                    <NavIcon>
                        <MdOutlineDashboard style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Dashboard</NavText>
                </NavItem>
                <NavItem eventKey="Deployment">
                    <NavIcon>
                        <GiServerRack style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Deployment</NavText>
                    <NavItem><NavText>All Devices</NavText></NavItem>
                    <NavItem><NavText>In-Progress Devices</NavText></NavItem>
                    <NavItem><NavText>Failed Devices</NavText></NavItem>
                    <NavItem><NavText>Recently Completed Devices</NavText></NavItem>
                </NavItem>
                <NavItem eventKey="cm">
                    <NavIcon>
                        <FaProjectDiagram style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Project Templates</NavText>
                    <NavItem><NavText>New Template</NavText></NavItem>
                    <NavItem><NavText>Edit Template</NavText></NavItem>
                    <NavItem><NavText>View Templates</NavText></NavItem>
                </NavItem>
                <NavItem eventKey="cm_tasks">
                    <NavIcon>
                        <TbSubtask style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Task Builder</NavText>
                    <NavItem><NavText>New Task</NavText></NavItem>
                    <NavItem><NavText>Edit Task</NavText></NavItem>
                    <NavItem><NavText>View Tasks</NavText></NavItem>
                </NavItem>
                <NavItem eventKey="troubleticket">
                    <NavIcon>
                        <FaBug style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Trouble Ticket</NavText>
                </NavItem>
                <NavItem eventKey="reports">
                    <NavIcon>
                        <HiOutlineDocumentReport style={{fontSize: "1.5em"}} />
                    </NavIcon>
                    <NavText>Reports</NavText>
                    <NavItem><NavText>Device Inventory</NavText></NavItem>
                    <NavItem><NavText>First Pass Yield</NavText></NavItem>
                    <NavItem><NavText>Failure Trends</NavText></NavItem>
                </NavItem>
            </SideNav.Nav>
        </SideNav>
  )
}

export default MySideNav