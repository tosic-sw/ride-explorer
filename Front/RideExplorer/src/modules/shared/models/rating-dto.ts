export interface ViewRatingDTO {
    id: number;
    evaluator: string;
    evaluated: string;
    driveId: number;
    positive: boolean;
    text: string;
}

export interface RatingDTO {
    evaluated: string;
    driveId: number;
    positive: boolean;
    text: string;
}

export interface RatingModalData {
    text: string;
    positive: boolean;
    username: string;
}