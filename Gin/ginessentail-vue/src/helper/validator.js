const telephoneValidator = (value) => /^1[3456789]\d{9}$/.test(value);

export default {
    telephoneValidator,
}