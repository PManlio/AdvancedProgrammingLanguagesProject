export interface InfoDiario {
    emailPaziente: string;
    text: string;
    sentiment: {
        polarity: number;
        subjectivity: number;
    };
    date: string;
}
