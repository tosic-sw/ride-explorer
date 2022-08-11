export interface ViewRatingDTO {
    id: number;
    evaluator: string;
    evaluated: string;
    driveId: number;
    positive: boolean;
    text: string;
}