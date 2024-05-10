import {
  DiatonicInterval,
  Interval,
  Mode,
  Note,
  Scale,
  VoiceLeadingChord,
  majorScale,
  majorScaleIntervals,
  scaleDegree,
  shiftChordTonesByMap,
  shiftNote,
  shiftScaleDiatonically,
  vlChordNotes,
} from "../src/chord_anthology"

describe("shiftNote", () => {

  it("should move a note by a fourth", () => {
    expect(shiftNote(Note.C, Interval.Fourth)).toEqual(Note.F)
  })

  it("should wrap around when shifting beyond a B", () => {
    expect(shiftNote(Note.A, Interval.MinorSeven)).toEqual(Note.G)
  })

  it("shifting by several intervals which make an octave gives the same note", () => {
    expect((() => {
      const n = Note.D
      const o = shiftNote(n, Interval.MajorThird)
      const p = shiftNote(o, Interval.MinorThird)
      const q = shiftNote(p, Interval.Fourth)
      return q
    })()).toEqual(Note.D)

    expect((() => {
      const n = Note.F
      const o = shiftNote(n, Interval.PerfectFifth)
      return shiftNote(o, Interval.Fourth)
    })()).toEqual(Note.F)
  })
})

describe("majorScaleIntervals", () => {
  it("should return correct intervals for all seven modes", () => {
    expect(majorScaleIntervals(Mode.Ionian)).toEqual([
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
    ])
    expect(majorScaleIntervals(Mode.Dorian)).toEqual([
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
    ])
    expect(majorScaleIntervals(Mode.Phrygian)).toEqual([
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
    ])
    expect(majorScaleIntervals(Mode.Lydian)).toEqual([
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
    ])
    expect(majorScaleIntervals(Mode.Mixolydian)).toEqual([
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
    ])
    expect(majorScaleIntervals(Mode.Aeolian)).toEqual([
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
    ])
    expect(majorScaleIntervals(Mode.Locrian)).toEqual([
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MinorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
      Interval.MajorSecond,
    ])
  })
})

describe("scaleDegree", () => {
  it("should return a degree", () => {
    const AMajorScale: Scale = {
      root: Note.A,
      intervals: majorScaleIntervals(Mode.Ionian),
    }
    expect(scaleDegree(AMajorScale, 1)).toEqual(Note.A)
    expect(scaleDegree(AMajorScale, 2)).toEqual(Note.B)
    expect(scaleDegree(AMajorScale, 3)).toEqual(Note.Db)
    expect(scaleDegree(AMajorScale, 4)).toEqual(Note.D)
    expect(scaleDegree(AMajorScale, 5)).toEqual(Note.E)
    expect(scaleDegree(AMajorScale, 7)).toEqual(Note.Ab)
    expect(scaleDegree(AMajorScale, 9)).toEqual(Note.B)
    expect(scaleDegree(AMajorScale, 11)).toEqual(Note.D)
    expect(scaleDegree(AMajorScale, 13)).toEqual(Note.Gb)

    const GbMajorScale: Scale = {
      root: Note.Gb,
      intervals: majorScaleIntervals(Mode.Ionian),
    }
    expect(scaleDegree(GbMajorScale, 2)).toEqual(Note.Ab)
    expect(scaleDegree(GbMajorScale, 9)).toEqual(Note.Ab)
  })
})

describe("voice leading chord", () => {

  it("should return its notes", () => {
    expect((() => {
      const cMaj: VoiceLeadingChord = {
        scale: majorScale(Note.C, Mode.Ionian),
        tones: [1,3,5],
      }
      return vlChordNotes(cMaj)
    })()).toEqual([Note.C, Note.E, Note.G])
  })

  it("should return its notes in order when its inverted", () => {
    expect((() => {
      const cMaj: VoiceLeadingChord = {
        scale: majorScale(Note.C, Mode.Ionian),
        tones: [3,1,5],
      }
      return vlChordNotes(cMaj)
    })()).toEqual([Note.E, Note.C, Note.G])

    expect((() => {
      const cMaj: VoiceLeadingChord = {
        scale: majorScale(Note.C, Mode.Ionian),
        tones: [3,5,1,7],
      }
      return vlChordNotes(cMaj)
    })()).toEqual([Note.E, Note.G, Note.C, Note.B])
  })
})

describe("shiftScaleDiatonically", () => {

  it("should change the scale of the chord by the step specified", () => {
    expect(
      shiftScaleDiatonically(
        majorScale(Note.C, Mode.Ionian),
        DiatonicInterval.Fourth,
      ),
    ).toEqual(
      majorScale(Note.F, Mode.Lydian),
    )

    expect(
      shiftScaleDiatonically(
        majorScale(Note.A, Mode.Aeolian),
        DiatonicInterval.Second,
      ),
    ).toEqual(
      majorScale(Note.B, Mode.Locrian),
    )

    expect(
      shiftScaleDiatonically(
        majorScale(Note.Gb, Mode.Lydian),
        DiatonicInterval.Seventh,
      ),
    ).toEqual(
      majorScale(Note.F, Mode.Phrygian),
    )

    expect(
      shiftScaleDiatonically(
        majorScale(Note.Eb, Mode.Mixolydian),
        DiatonicInterval.Sixth,
      ),
    ).toEqual(
      majorScale(Note.C, Mode.Phrygian),
    )
  })
})

describe("shiftChordTonesByMap", () => {

  it("throws error when inversion map not specified", () => {
    const eMaj: VoiceLeadingChord = {
      scale: majorScale(Note.Gb, Mode.Ionian),
      tones: [5,1],
    }
    expect(() => {shiftChordTonesByMap(eMaj, new Map<number, number>())}).toThrow()
  })

  it("throws error when map does not have same number of entries as chord tones", () => {
    const eMaj: VoiceLeadingChord = {
      scale: majorScale(Note.Gb, Mode.Ionian),
      tones: [5,1,3],
    }
    expect(() => {
      shiftChordTonesByMap(
        eMaj,
        new Map<number, number>(
          [
            [1,1],
            [1,1],
          ],
        ),
      )
    }).toThrow()
    expect(() => {
      shiftChordTonesByMap(
        eMaj,
        new Map<number, number>(
          [
            [1,1],
            [1,1],
            [1,1],
            [1,1],
          ],
        ),
      )
    }).toThrow()
  })

  it("changes the chord tones according to the map", () => {
    expect(shiftChordTonesByMap(
      {
        scale: majorScale(Note.Gb, Mode.Ionian),
        tones: [5,1,3],
      },
      new Map<number, number>(
        [
          [5,1],
          [1,3],
          [3,5],
        ],
      ),
    )).toEqual({
      scale: majorScale(Note.Gb, Mode.Ionian),
      tones: [1,3,5],
    })

    expect(shiftChordTonesByMap(
      {
        scale: majorScale(Note.Bb, Mode.Ionian),
        tones: [5,1,3],
      },
      new Map<number, number>(
        [
          [5,6],
          [1,2],
          [3,4],
        ],
      ),
    )).toEqual({
      scale: majorScale(Note.Bb, Mode.Ionian),
      tones: [6,2,4],
    })
  })
})

describe("POC creating cycle2", () => {

  it("logs cycle 2 chord notes - triads; close voice", () => {

    let c: VoiceLeadingChord = {
      scale: majorScale(Note.C, Mode.Ionian),
      tones: [1, 5, 3],
    }

    for (let i=0; i < 10; i++) {

      console.log(vlChordNotes(c))

      c = shiftChordTonesByMap(c, new Map<number, number>(
        [
          [1, 5],
          [3, 1],
          [5, 3],
        ],
      ))

      c.scale = shiftScaleDiatonically(c.scale, DiatonicInterval.Second)
    }

  })

})
