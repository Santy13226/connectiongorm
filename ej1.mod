BEGIN;
CREATE: EX(1, 3): MARK(1);
FARMAC ASSIGN: A(2)=1;
    QUEUE,1,3,APREND;
    SEIZE: FARMACEUTICO;
    DELAY: EX(2,3);
    RELEASE: FARMACEUTICO:NEXT(CAJER);
APREND QUEUE,2;
    SEIZE: APRENDIZ;
    DELAY: EX(2,4);
    RELEASE: APRENDIZ;
CAJER QUEUE,3;
    SEIZE:CAJERO;
    DELAY: EX(3,1);
    RELEASE: CAJERO;
COUNT: A(2);
TALLY: A(2),INT(1):DISPOSE;
END;