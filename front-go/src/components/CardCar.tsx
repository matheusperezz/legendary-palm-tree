import React, { useState } from "react";
import styled from 'styled-components';
import axios from "axios";

const StyledCard = styled.div`

    text-align: center;

    h3 {
        margin: 25px;
        padding: 0 0 5px 0;
        font-size: 16px;
    }

    p {
        margin: 25px;
        color: #666c74;
        font-size: 18px;
        line-height: 20px;
    }

`;

interface Car {
    id: number,
    marca: string,
    modelo: string,
    motor: string,
    ano: string,
    potencia_em_cavalos: string
}

function Carros() {

    const [cars, setCars] = useState<Car[]>([]);

    axios.get('http://localhost:8000/api/carros')
        .then(res => {
            setCars(res.data);
        });

    return (
        <StyledCard>
            {cars.map(car => (
                <React.Fragment key={car.id}>
                    <h3>{car.marca} | {car.modelo}</h3>
                    <p>Motor: {car.motor}</p>
                    <p>Ano: {car.ano}</p>
                    <p>HP: {car.potencia_em_cavalos}</p>
                </React.Fragment>
            ))}
        </StyledCard>
    )
}

export default Carros;