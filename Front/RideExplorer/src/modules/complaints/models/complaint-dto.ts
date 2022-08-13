export interface ComplaintDTO {
    accuser: string
    accused: string
    driveId: number
    text: string
    createdAt: number
    id: number
}

export interface CreateComplaintDTO {
    accused: string
    driveId: number
    text: string
}