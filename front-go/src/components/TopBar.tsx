import React from 'react';
import logo_go from '../assets/logo-go.png';
import styled from 'styled-components';

const StyledTopBar = styled.header`
    
    display: flex;
    align-items: center;
    width: 100%;
    max-width: 100vw;
    height: auto;
    min-height: 5vh;
    flex-direction: row;
    background-color: #282c34;
    justify-content: space-between;

    h1 {
        width: auto;
        font-size: 32px;
        margin: 8px;
        color: #e8e8e8;
    }

    img {
        height: 8vmin;
        pointer-events: none;
        animation: rotate 10s linear infinite;
        margin: 8px;
    }

    button {
        padding: 12px;
        margin-right: 16px;
        font-size: 14px;
        font-weight: 900;
        border-radius: 10%;
        color: #383838;
        background-color: #b7b7b7;
        transition: background-color 0.5s ease;
        border: none;
    }

    button:hover {
        background-color: #393939;
        color: #b7b7b7;
        cursor: pointer;
    }

    @keyframes rotate {
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);
        }
    }
    
`;

function TopBar() {
    return (
        <StyledTopBar>
            <img src={logo_go} alt="" />
            <h1>API Rest Carros</h1>
            <button>GitHub</button>
        </StyledTopBar>
    );
}

export default TopBar;