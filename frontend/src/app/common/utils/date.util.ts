export function toInputDateValue(value: Date | string): string {
    const date = (value instanceof Date) ? value : new Date(value);
    const yyyy = date.getFullYear();
    const mm = date.getUTCMonth() + 1;
    const dd = date.getUTCDate();
    return `${String(yyyy).padStart(4, '0')}-${String(mm).padStart(2, '0')}-${String(dd).padStart(2, '0')}`;
}
